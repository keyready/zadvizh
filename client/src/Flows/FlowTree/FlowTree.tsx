import { Background, type Edge, type Node, ReactFlow } from '@xyflow/react';

import '@xyflow/react/dist/style.css';
import { RootGroup } from '../FlowNode/RootGroup.tsx';
import { SubrootGroup } from '../FlowNode/SubrootGroup.tsx';

const nodeTypes = {
    RootGroup: RootGroup,
    SubrootGroup: SubrootGroup,
};

const edges: Edge[] = [
    { id: 'b-c', source: 'B', target: 'C' },
    { id: 'dev-dev', source: 'A', sourceHandle: 's-dev', target: 'dev' },
    { id: 'sec-sec', source: 'A', sourceHandle: 's-sec', target: 'sec' },
    { id: 'devops-devops', source: 'A', sourceHandle: 's-devops', target: 'devops' },
    { id: 'science-science', source: 'A', sourceHandle: 's-science', target: 'science' },
];

const nodes: Node[] = [
    {
        id: 'A',
        type: 'RootGroup',
        data: { label: null },
        position: { x: 0, y: 0 },
    },
    {
        id: 'dev',
        data: { label: 'Разработка' },
        position: { x: -150, y: 200 },
        style: {
            color: 'black',
        },
    },
    {
        id: 'sec',
        data: { label: 'Безопасность' },
        position: { x: 0, y: 200 },
        style: {
            color: 'black',
        },
    },
    {
        id: 'devops',
        data: { label: 'Безопасность' },
        position: { x: 150, y: 200 },
        style: {
            color: 'black',
        },
    },
    {
        id: 'science',
        data: { label: 'Безопасность' },
        position: { x: 300, y: 200 },
        style: {
            color: 'black',
        },
    },
    {
        id: 'B',
        type: 'input',
        data: { label: 'Бирюков Денис' },
        position: { x: 125, y: 15 },
        parentId: 'A',
        extent: 'parent',
        style: {
            width: 150,
            height: 50,
            color: 'black',
        },
    },
    {
        id: 'C',
        data: { label: 'Дудкин Андрей' },
        position: { x: 125, y: 85 },
        parentId: 'A',
        extent: 'parent',
        style: {
            width: 150,
            height: 50,
            color: 'black',
        },
    },
];

export const FlowTree = () => {
    return (
        <div className="bg-main-gradient h-screen w-full">
            <ReactFlow edges={edges} nodes={nodes} nodeTypes={nodeTypes}>
                <Background />
            </ReactFlow>
        </div>
    );
};
