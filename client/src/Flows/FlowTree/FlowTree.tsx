import { Background, type Edge, MiniMap, type Node, ReactFlow, useReactFlow } from '@xyflow/react';
import '@xyflow/react/dist/style.css';
import dagre from '@dagrejs/dagre';
import { useCallback, useEffect, useMemo, useState } from 'react';
import { Spinner } from '@heroui/react';

import { RootGroup } from '../FlowNode/RootGroup.tsx';
import { SubrootGroup } from '../FlowNode/SubrootGroup.tsx';

import { SourceNodesMap, transformData } from './lib/generateNodes.ts';
import { LeafDrawer } from './LeafDrawer.tsx';
import { treeSearch } from './lib/treeSearch.ts';
import { FlowTreeSearchPanel } from './FlowTreeSearchPanel.tsx';

const dagreGraph = new dagre.graphlib.Graph().setDefaultEdgeLabel(() => ({}));

const nodeWidth = 70;
const nodeHeight = 70;

const nodeTypes = {
    RootGroup: RootGroup,
    SubrootGroup: SubrootGroup,
};

export const FlowTree = () => {
    // const navigate = useNavigate();

    const [hierarchy, setHierarchy] = useState<SourceNodesMap[]>([]);
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [searchValue, setSearchValue] = useState<string>('');

    useEffect(() => {
        const getEmployers = async () => {
            try {
                const result = await fetch('/api/v1/employers', {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': localStorage.getItem('t') || '',
                    },
                });
                setHierarchy([await result.json()]);
            } catch (e) {
                // navigate('/');
            } finally {
                setIsLoading(false);
            }
        };

        setIsLoading(true);
        getEmployers();
    }, []);

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

    const { edges: formattedEdges, nodes: formattedNodes } = transformData(hierarchy || []);
    const { nodes: layoutedNodes, edges: layoutedEdges } = getLayoutedElements(
        formattedNodes,
        formattedEdges,
    );

    const { setCenter } = useReactFlow();

    const [searchResults, setSearchResults] = useState<Node[]>([]);

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

    useEffect(() => {
        setSearchResults(treeSearch(layoutedNodes, searchValue));
    }, [searchValue]);

    const [selectedNode, setSelectedNode] = useState<SourceNodesMap | undefined>(undefined);

    const handleNodeClick = useCallback(async (_: any, node: Node) => {
        setSelectedNode(node as unknown as SourceNodesMap);
        await setCenter(node.position.x + 70, node.position.y + 70, { zoom: 3, duration: 500 });
    }, []);

    if (isLoading) {
        return (
            <div className="bg-main-gradient relative h-screen w-full">
                <div className="absolute bottom-0 left-0 right-0 top-0 flex items-center justify-center bg-primary bg-opacity-50 backdrop-blur">
                    <div className="flex h-32 w-64 items-center justify-center rounded-md bg-gray-400 bg-opacity-50">
                        <Spinner size="lg" />
                    </div>
                </div>
            </div>
        );
    }

    return (
        <div className="bg-main-gradient relative h-screen w-full">
            <FlowTreeSearchPanel
                searchResults={searchResults}
                searchValue={searchValue}
                setSearchValue={setSearchValue}
                handleNodeClick={handleNodeClick}
            />

            <ReactFlow
                onNodeClick={handleNodeClick}
                edges={edges}
                nodes={nodes}
                nodeTypes={nodeTypes}
                fitView
            >
                <Background />
                <MiniMap bgColor="gray" maskColor="#444" />
            </ReactFlow>

            <LeafDrawer selectedNode={selectedNode} setSelectedNode={setSelectedNode} />
        </div>
    );
};
