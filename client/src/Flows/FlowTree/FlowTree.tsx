import { Background, type Edge, MiniMap, type Node, ReactFlow } from '@xyflow/react';
import '@xyflow/react/dist/style.css';
import dagre from '@dagrejs/dagre';
import { useCallback, useEffect, useMemo, useState } from 'react';

import { RootGroup } from '../FlowNode/RootGroup.tsx';
import { SubrootGroup } from '../FlowNode/SubrootGroup.tsx';

import { rawData, SourceNodesMap, transformData } from './lib/generateNodes.ts';
import { LeafDrawer } from './LeafDrawer.tsx';

const dagreGraph = new dagre.graphlib.Graph().setDefaultEdgeLabel(() => ({}));

const nodeWidth = 70;
const nodeHeight = 70;

const nodeTypes = {
    RootGroup: RootGroup,
    SubrootGroup: SubrootGroup,
};

export const FlowTree = () => {
    const getLayoutedElements = useCallback(
        (nodes: Node[], edges: Edge[], direction = 'TB'): { nodes: Node[]; edges: Edge[] } => {
            dagreGraph.setGraph({ rankdir: direction, nodesep: 125 });

            nodes.forEach((node) => {
                dagreGraph.setNode(node.id, { width: nodeWidth, height: nodeHeight });
            });

            edges.forEach((edge) => {
                dagreGraph.setEdge(edge.source, edge.target);
            });

            dagre.layout(dagreGraph);

            const newNodes = nodes.map((node) => {
                const nodeWithPosition = dagreGraph.node(node.id);
                return {
                    ...node,
                    targetPosition: 'top',
                    sourcePosition: 'bottom',
                    position: {
                        x: nodeWithPosition.x - nodeWidth / 2,
                        y: nodeWithPosition.y - nodeHeight / 2,
                    },
                };
            });

            // @ts-expect-error несоответсвие указанных типов
            return { nodes: newNodes, edges };
        },
        [],
    );

    const { edges: formattedEdges, nodes: formattedNodes } = transformData(rawData || []);

    const { nodes: layoutedNodes, edges: layoutedEdges } = getLayoutedElements(
        formattedNodes,
        formattedEdges,
    );

    const nodes = useMemo(() => layoutedNodes, [layoutedNodes]);
    const edges = useMemo(() => layoutedEdges, [layoutedEdges]);

    useEffect(() => {
        const banner = document.getElementsByClassName(
            'react-flow__panel react-flow__attribution bottom right',
        );
        if (banner.length) {
            banner[0].remove();
        }
    }, []);

    const [selectedNode, setSelectedNode] = useState<SourceNodesMap | undefined>(undefined);

    const handleNodeClick = useCallback((_: any, node: Node) => {
        setSelectedNode(node as unknown as SourceNodesMap);
    }, []);

    return (
        <div className="bg-main-gradient h-screen w-full">
            <ReactFlow
                onNodeClick={handleNodeClick}
                edges={edges}
                nodes={nodes}
                nodeTypes={nodeTypes}
            >
                <Background />
                <MiniMap bgColor="gray" maskColor="#444" />
            </ReactFlow>

            <LeafDrawer selectedNode={selectedNode} setSelectedNode={setSelectedNode} />
        </div>
    );
};
