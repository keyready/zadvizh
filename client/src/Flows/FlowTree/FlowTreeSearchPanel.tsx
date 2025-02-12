import { Input } from '@heroui/input';
import { Divider, Kbd } from '@heroui/react';
import { RiSearchLine } from '@remixicon/react';
import { type Node } from '@xyflow/react';
import { useEffect, useRef, useState } from 'react';

interface FlowTreeSearchPanelProps {
    searchResults: Node[];
    searchValue: string;
    setSearchValue: (str: string) => void;
    handleNodeClick: (event: any, node: Node) => void;
}

export const FlowTreeSearchPanel = (props: FlowTreeSearchPanelProps) => {
    const { searchResults, searchValue, setSearchValue, handleNodeClick } = props;
    const [focusedIndex, setFocusedIndex] = useState<number | null>(null);

    const resultRefs = useRef<(HTMLButtonElement | null)[]>([]);
    const inputRef = useRef<HTMLInputElement | null>(null);

    const handleKeyDown = (ev: React.KeyboardEvent<HTMLInputElement> | KeyboardEvent) => {
        if (ev.code === 'KeyK' && ev.ctrlKey) {
            ev.preventDefault();
            if (inputRef.current) {
                inputRef.current.focus();
            }
        }

        if (!searchValue || !searchResults.length) return;

        if (ev.key === 'Escape') {
            setSearchValue('');
        } else if (ev.key === 'ArrowDown') {
            ev.preventDefault();
            setFocusedIndex((prevIndex) =>
                prevIndex !== null && prevIndex < searchResults.length - 1 ? prevIndex + 1 : 0,
            );
        } else if (ev.key === 'ArrowUp') {
            ev.preventDefault();
            setFocusedIndex((prevIndex) =>
                prevIndex !== null && prevIndex > 0 ? prevIndex - 1 : searchResults.length - 1,
            );
        } else if (ev.key === 'Enter' && focusedIndex !== null) {
            ev.preventDefault();
            const selectedNode = searchResults[focusedIndex];
            setSearchValue('');
            handleNodeClick({}, selectedNode);
            setFocusedIndex(null);
        }
    };

    useEffect(() => {
        document.addEventListener('keydown', handleKeyDown);

        return () => {
            document.removeEventListener('keydown', handleKeyDown);
        };
    }, [searchValue, searchResults, focusedIndex, setSearchValue, handleNodeClick]);

    useEffect(() => {
        if (focusedIndex !== null) {
            resultRefs.current[focusedIndex]?.focus();
        }
    }, [focusedIndex]);

    return (
        <div className="absolute left-2 top-5 z-50 max-h-72 w-2/6 rounded-lg bg-primary bg-opacity-40 p-4">
            <div className="relative flex items-center justify-start gap-4">
                <Input
                    startContent={<RiSearchLine className="opacity-40" />}
                    endContent={
                        <Kbd className="absolute right-5 z-50" keys={['ctrl']}>
                            K
                        </Kbd>
                    }
                    ref={inputRef}
                    onKeyDown={handleKeyDown}
                    value={searchValue}
                    onValueChange={setSearchValue}
                    label="Поиск по коллективу"
                    classNames={{
                        inputWrapper:
                            'bg-primary bg-opacity-50 ' +
                            'data-[hover=true]:bg-primary data-[hover=true]:bg-opacity-80 ' +
                            'group-data-[focus=true]:bg-primary',
                    }}
                />
                {/*{searchValue && (*/}
                {/*    <Dropdown>*/}
                {/*        <DropdownTrigger>*/}
                {/*            <Button*/}
                {/*                className={*/}
                {/*                    'w-64 bg-primary bg-opacity-50 ' +*/}
                {/*                    'data-[hover=true]:bg-primary data-[hover=true]:bg-opacity-80' +*/}
                {/*                    'group-data-[focus=true]:bg-primary'*/}
                {/*                }*/}
                {/*            >*/}
                {/*                Искать по...*/}
                {/*            </Button>*/}
                {/*        </DropdownTrigger>*/}
                {/*        <DropdownMenu color="primary">*/}
                {/*            <DropdownItem key="label">Названию</DropdownItem>*/}
                {/*            <DropdownItem key="label">Специальности</DropdownItem>*/}
                {/*            <DropdownItem key="label">Названию</DropdownItem>*/}
                {/*        </DropdownMenu>*/}
                {/*    </Dropdown>*/}
                {/*)}*/}
            </div>
            {searchValue && searchResults?.length ? (
                <>
                    <Divider className="my-3 w-full" />
                    <div className="flex h-32 flex-col items-start gap-0.5 overflow-auto">
                        {searchResults.map((sr, index) => (
                            <button
                                // @ts-expect-error non types
                                ref={(el) => (resultRefs.current[index] = el)}
                                className={
                                    'w-full rounded-md bg-primary' +
                                    ' bg-opacity-50 p-2 text-start' +
                                    ' focus-visible:outline-none' +
                                    (focusedIndex === index ? ' bg-primary-100 bg-opacity-100' : '')
                                }
                                type="button"
                                onClick={() => {
                                    setSearchValue('');
                                    handleNodeClick({}, sr);
                                    setFocusedIndex(null);
                                }}
                                key={sr.id}
                            >
                                {sr.data.label as string}
                            </button>
                        ))}
                    </div>
                </>
            ) : searchValue ? (
                <p className="mt-2 italic text-danger">Результатов не найдено</p>
            ) : null}
        </div>
    );
};
