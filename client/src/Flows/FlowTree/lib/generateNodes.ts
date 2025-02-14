import { type Edge, type Node } from '@xyflow/react';

interface NodeData {
    'label': string;
    'data-label'?: string;
    'teamRole'?: string;
    'position'?: string;
    'children'?: SourceNodesMap[];
}

export interface SourceNodesMap {
    id: string;
    data: NodeData;
    children?: SourceNodesMap[];
    [key: string]: any;
}

export function transformData(input: any) {
    let nodeId = 1;
    const nodes: Node[] = [];
    const edges: Edge[] = [];

    function getNodeColors(dataLabel: string): { border: string; bg: string } {
        switch (dataLabel) {
            case 'org': {
                return { border: '2px solid #95bf74', bg: 'rgba(149,191,116,0.75)' };
            }
            case 'dev': {
                return { border: '2px solid #258ea6', bg: 'rgba(37,142,166,0.5)' };
            }
            case 'sec': {
                return { border: '2px solid #6f1d1b', bg: 'rgba(111,29,27,0.5)' };
            }
            case 'devops': {
                return { border: '2px solid #dacc3e', bg: 'rgba(218,204,62,0.5)' };
            }
            case 'science': {
                return { border: '2px solid #e2c2ff', bg: 'rgba(226,194,255,0.5)' };
            }
            default: {
                return { border: '2px solid #fefefe', bg: '#F5212111' };
            }
        }
    }

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
                ...node.data,
                children: node.children,
            },
            position: { x: 0, y: 0 },
            style: {
                border: getNodeColors(node.data['data-label'] || '').border,
                background: getNodeColors(node.data['data-label'] || '').bg,
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
                    type: 'default',
                    style: {
                        strokeWidth: 1,
                        stroke: '#95bf74',
                    },
                });
            });
        }

        return currentNodeId;
    }

    input.forEach((root: SourceNodesMap) => traverseTree(root));

    return { nodes, edges };
}

export const rawData: SourceNodesMap[] = [
    {
        id: 'sgsdfgsdfgsdfg',
        data: { 'label': 'Бирюков Д.Н.', 'data-label': 'org' },
        children: [
            {
                id: 'sgsdfgsdfg234234sdfg',
                data: { 'label': 'Дудкин А.С.', 'data-label': 'org' },
                children: [
                    {
                        id: '134',
                        data: { 'label': 'Разработка', 'data-label': 'dev' },
                        children: [
                            {
                                id: 'dev-team-1',
                                data: { label: 'NexusX Team' },
                                children: [
                                    {
                                        id: 'NexusX-Team-1',
                                        data: {
                                            label: 'Корчак Р.Д.',
                                            teamRole: 'cap',
                                            position: 'Frontend',
                                        },
                                    },
                                    {
                                        id: 'NexusX-Team-2',
                                        data: {
                                            label: 'Кофанов В.С.',
                                            teamRole: 'part',
                                            position: 'Backend',
                                        },
                                    },
                                    {
                                        id: 'NexusX-Team-3',
                                        data: {
                                            label: 'Поляков Д.С.',
                                            teamRole: 'part',
                                            position: 'UI/UX',
                                        },
                                    },
                                    {
                                        id: 'NexusX-Team-4',
                                        data: {
                                            label: 'Терещенко А.А.',
                                            teamRole: 'part',
                                            position: 'Mobile Dev',
                                        },
                                    },
                                ],
                            },
                            {
                                id: 'dev-team-2',
                                data: { label: 'Konyhov&Co' },
                                children: [
                                    {
                                        id: 'Konyhov&Co-1',
                                        data: {
                                            label: 'Коныхов В.С.',
                                            teamRole: 'part',
                                            position: 'Fullstack',
                                        },
                                    },
                                    {
                                        id: 'Konyhov&Co-2',
                                        data: {
                                            label: 'Тараскин Т.Д.',
                                            teamRole: 'cap',
                                            position: 'Fullstack',
                                        },
                                    },
                                    {
                                        id: 'Konyhov&Co-3',
                                        data: {
                                            label: 'Ильчук Д.Д.',
                                            teamRole: 'part',
                                            position: 'Product Manager',
                                        },
                                    },
                                ],
                            },
                        ],
                    },
                    {
                        id: '1341234',
                        data: { 'label': 'Безопасность', 'data-label': 'sec' },
                        children: [
                            {
                                id: 'sec-team-1',
                                data: { label: 'Red Cadets' },
                                children: [
                                    { id: 'Red-Cadets-1', data: { label: 'Яроцкий Г.Д.' } },
                                    { id: 'Red-Cadets-2', data: { label: 'Тыренко Д.В.' } },
                                    { id: 'Red-Cadets-3', data: { label: 'Степанов А.С.' } },
                                ],
                            },
                        ],
                    },
                    {
                        id: '13342344',
                        data: { 'label': 'DevOps', 'data-label': 'devops' },
                        children: [
                            {
                                id: 'devops-1',
                                data: { label: 'Крюков Р.О.' },
                                children: [
                                    {
                                        id: 'devops-1-1',
                                        data: { label: 'Шаповалов Д.Г.' },
                                        children: [],
                                    },
                                    {
                                        id: 'devops-1-1',
                                        data: { label: 'Косарев А.А.' },
                                        children: [],
                                    },
                                ],
                            },
                        ],
                    },
                    {
                        id: '134',
                        data: { 'label': 'Научная деятельность', 'data-label': 'science' },
                        children: [
                            {
                                id: 'science-1',
                                data: { label: 'Мишуков О.В.' },
                                children: [
                                    {
                                        id: 'science-1-1',
                                        data: { label: 'Участник научки 1' },
                                        children: [],
                                    },
                                    {
                                        id: 'science-1-2',
                                        data: { label: 'Участник научки 2' },
                                        children: [],
                                    },
                                ],
                            },
                        ],
                    },
                ],
            },
        ],
    },
];
