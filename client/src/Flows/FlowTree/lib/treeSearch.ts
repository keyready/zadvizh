export function treeSearch(tree, predicate) {
    let results = [];

    // Если текущий узел удовлетворяет предикату, добавляем его в результат
    if (predicate(tree) && !tree?.children?.length) {
        results.push(tree);
    }

    // Если у текущего узла есть дочерние элементы, рекурсивно ищем в них
    if (tree.children && tree.children.length > 0) {
        for (const child of tree.children) {
            results = results.concat(treeSearch(child, predicate));
        }
    }

    return results;
}
