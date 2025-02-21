import { Input } from '@heroui/input';
import {
    Button,
    Divider,
    Dropdown,
    DropdownItem,
    DropdownMenu,
    DropdownTrigger,
    Kbd,
} from '@heroui/react';
import { RiSearchLine } from '@remixicon/react';
import { type Node } from '@xyflow/react';
import { useEffect, useMemo, useRef, useState } from 'react';

enum SearchFieldTypes {
    Label = 'label',
    Field = 'field',
}

const SearchFieldMapper: Record<SearchFieldTypes, string> = {
    [SearchFieldTypes.Field]: 'Род деятельности',
    [SearchFieldTypes.Label]: 'Фамилия',
};

interface FlowTreeSearchPanelProps {
    searchResults: Node[];
    searchValue: string;
    setSearchValue: (str: string) => void;
    handleNodeClick: (event: any, node: Node) => void;
    setSearchField: (val: string) => void;
}

export const FlowTreeSearchPanel = (props: FlowTreeSearchPanelProps) => {
    const { searchResults, searchValue, setSearchField, setSearchValue, handleNodeClick } = props;

    const [focusedIndex, setFocusedIndex] = useState<number | null>(null);

    const resultRefs = useRef<(HTMLButtonElement | null)[]>([]);
    const inputRef = useRef<HTMLInputElement | null>(null);

    const [selectedSearchFields, setSelectedSearchFields] = useState(new Set(['label']));

    const selectedSearchField = useMemo(
        () => Array.from(selectedSearchFields).join(', ').replace(/_/g, '') as SearchFieldTypes,
        [selectedSearchFields],
    );

    useEffect(() => {
        setSearchField(selectedSearchField);
    }, [selectedSearchField]);

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
        <div className="absolute left-2 top-16 z-50 max-h-96 rounded-lg bg-primary bg-opacity-40 p-4 lg:w-2/6">
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
                {searchValue && (
                    <Dropdown>
                        <DropdownTrigger>
                            <Button
                                className={
                                    'bg-primary bg-opacity-50 lg:w-64 ' +
                                    'data-[hover=true]:bg-primary data-[hover=true]:bg-opacity-80' +
                                    'group-data-[focus=true]:bg-primary'
                                }
                            >
                                {SearchFieldMapper[selectedSearchField] || 'Искать по...'}
                            </Button>
                        </DropdownTrigger>
                        <DropdownMenu
                            selectedKeys={selectedSearchFields}
                            // @ts-expect-error types mismatch old version
                            onSelectionChange={setSelectedSearchFields}
                            selectionMode="single"
                            color="primary"
                        >
                            <DropdownItem key="label">Фамилии</DropdownItem>
                            <DropdownItem key="position">Род деятельности</DropdownItem>
                        </DropdownMenu>
                    </Dropdown>
                )}
            </div>
            {searchValue && searchResults?.length ? (
                <>
                    <Divider className="my-3 w-full" />
                    <div className="flex h-64 flex-col items-start gap-0.5 overflow-auto">
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
                                <p>{sr.data.label as string}</p>
                                <p className="text-xs italic opacity-50">
                                    {sr.data?.position as string}
                                </p>
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
