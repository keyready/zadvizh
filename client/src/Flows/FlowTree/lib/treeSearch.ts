import { type Node } from '@xyflow/react';

export function treeSearch(tree: Node[], search: string, searchField: string = 'label'): Node[] {
    const results: Node[] = [];

    tree.forEach((node) => {
        // @ts-expect-error non types
        if (node.data[searchField].toLowerCase().includes(search.toLowerCase())) {
            results.push(node);
        }
    });

    return results;
}
