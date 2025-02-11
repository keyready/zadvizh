import { SourceNodesMap } from './generateNodes.ts';

export function treeSearch(
    tree: SourceNodesMap,
    predicate: (node: SourceNodesMap) => boolean,
): SourceNodesMap[] {
    let results = [];

    if (predicate(tree) && !tree?.children?.length) {
        results.push(tree);
    }

    if (tree.children && tree.children.length > 0) {
        for (const child of tree.children) {
            results = results.concat(treeSearch(child, predicate));
        }
    }

    return results;
}
