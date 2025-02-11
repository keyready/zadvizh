import { type Edge, type Node } from '@xyflow/react';

export interface SourceNodesMap {
    name: string;
    children: SourceNodesMap[];
    [key: string]: any;
}

export function transformData(input: any) {
    let nodeId = 1;
    const nodes: Node[] = [];
    const edges: Edge[] = [];

    function traverseTree(node: SourceNodesMap, parentId: number | null = null, level: number = 0) {
        const currentNodeId = (nodeId += 1);

        let type = '';
        if (parentId === null) {
            type = 'input';
        } else if (node.children && node.children.length > 0) {
            type = '';
        } else {
            type = 'output';
        }

        nodes.push({
            id: currentNodeId.toString(),
            data: {
                label: node.name,
                name: `${node.firstname} ${node.middlename} ${node.lastname}`,
                address: 'node.address',
                department: 'node.department',
                division: 'node.division',
                position: 'node.position',
                phone: 'node.phone',
            },
            position: { x: 0, y: 0 },
            style: {
                border: '2px solid #F52121',
                background: '#F5212111',
                borderRadius: '12px',
            },
            type,
        });

        if (node.children && node.children.length > 0) {
            node.children.forEach((child) => {
                const childNodeId = traverseTree(child, currentNodeId, level + 1);
                edges.push({
                    id: `e${currentNodeId}${childNodeId}`,
                    source: currentNodeId.toString(),
                    target: childNodeId.toString(),
                    type: 'step',
                    style: {
                        strokeWidth: 1,
                        stroke: '#F5212144',
                    },
                });
            });
        }

        return currentNodeId;
    }

    input.forEach((root: SourceNodesMap) => traverseTree(root));

    return { nodes, edges };
}

export const rawData = [
    {
        name: 'AI Robotics Solutions',
        children: [
            {
                name: 'AI Development',
                children: [
                    {
                        name: 'Machine Learning',
                        children: [
                            {
                                name: 'Team A',
                                children: [{ name: 'Alice' }, { name: 'Bob' }, { name: 'Charlie' }],
                            },
                            {
                                name: 'Team B',
                                children: [
                                    { name: 'David' },
                                    { name: 'Eve' },
                                    { name: 'Frank' },
                                    { name: 'Grace' },
                                ],
                            },
                        ],
                    },
                    {
                        name: 'Computer Vision',
                        children: [
                            {
                                name: 'Team C',
                                children: [
                                    { name: 'Heidi' },
                                    { name: 'Ivan' },
                                    { name: 'Jack' },
                                    { name: 'Karen' },
                                ],
                            },
                        ],
                    },
                ],
            },
            {
                name: 'Robotics',
                children: [
                    {
                        name: 'Mechanical Engineering',
                        children: [
                            {
                                name: 'Team D',
                                children: [
                                    { name: 'Liam' },
                                    { name: 'Mia' },
                                    { name: 'Noah' },
                                    { name: 'Olivia' },
                                ],
                            },
                        ],
                    },
                    {
                        name: 'Automation',
                        children: [
                            {
                                name: 'Team E',
                                children: [{ name: 'Paul' }, { name: 'Quinn' }, { name: 'Rachel' }],
                            },
                            {
                                name: 'Team F',
                                children: [
                                    { name: 'Sophia' },
                                    { name: 'Tom' },
                                    { name: 'Uma' },
                                    { name: 'Victor' },
                                    { name: 'Wendy' },
                                ],
                            },
                        ],
                    },
                ],
            },
            {
                name: 'R&D',
                children: [
                    {
                        name: 'AI Research',
                        children: [
                            {
                                name: 'Team G',
                                children: [
                                    { name: 'Xander' },
                                    { name: 'Yara' },
                                    { name: 'Zane' },
                                    { name: 'Ava' },
                                ],
                            },
                        ],
                    },
                    {
                        name: 'Prototype Development',
                        children: [
                            {
                                name: 'Team H',
                                children: [
                                    { name: 'Benjamin' },
                                    { name: 'Clara' },
                                    { name: 'Dylan' },
                                    { name: 'Eva' },
                                    { name: 'Felix' },
                                ],
                            },
                        ],
                    },
                ],
            },
            {
                name: 'Sales & Marketing',
                children: [
                    {
                        name: 'Product Marketing',
                        children: [
                            {
                                name: 'Team I',
                                children: [
                                    { name: 'George' },
                                    { name: 'Hannah' },
                                    { name: 'Isabella' },
                                    { name: 'Jackie' },
                                ],
                            },
                        ],
                    },
                    {
                        name: 'Sales Strategy',
                        children: [
                            {
                                name: 'Team J',
                                children: [
                                    { name: 'Kevin' },
                                    { name: 'Lily' },
                                    { name: 'Monica' },
                                    { name: 'Nathan' },
                                    { name: 'Olga' },
                                ],
                            },
                        ],
                    },
                ],
            },
        ],
    },
];
