import { Divider, Drawer, DrawerContent } from '@heroui/react';
import { useCallback } from 'react';

import { SourceNodesMap } from './lib/generateNodes.ts';

interface LeafDrawerProps {
    selectedNode: SourceNodesMap | undefined;
    setSelectedNode: (node: SourceNodesMap | undefined) => void;
}

export const LeafDrawer = ({ selectedNode, setSelectedNode }: LeafDrawerProps) => {
    const generateTeamRoleLabel = useCallback((teamRole: string) => {
        switch (teamRole) {
            case 'cap':
                return 'Капитан';
            case 'part':
                return 'Участник';
            default:
                return '';
        }
    }, []);

    return (
        <Drawer size="sm" isOpen={!!selectedNode} onClose={() => setSelectedNode(undefined)}>
            <DrawerContent className="flex flex-col gap-2 bg-opacity-80 p-4">
                <h1 className="text-2xl">{selectedNode?.data?.label as string}</h1>

                {selectedNode?.data?.teamRole ? (
                    <>
                        <Divider className="h-0.5 w-full rounded-md" />
                        <h2 className="text-xl opacity-50">
                            Роль в команде:{' '}
                            {generateTeamRoleLabel(selectedNode?.data?.teamRole as string)}
                        </h2>
                        <h2 className="text-xl opacity-50">
                            Деятельность: {selectedNode?.data?.position as string}
                        </h2>
                    </>
                ) : null}

                {selectedNode?.data?.children?.length ? (
                    <>
                        <Divider className="h-0.5 w-full rounded-md" />
                        <h3 className="text-xl opacity-50">Участники:</h3>
                        {selectedNode?.data?.children?.map((ch: any) => (
                            <li key={ch.id}>{ch.data.label}</li>
                        ))}
                    </>
                ) : null}
            </DrawerContent>
        </Drawer>
    );
};
