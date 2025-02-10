import { Handle, Position } from '@xyflow/react';
import '@xyflow/react/dist/style.css';

export const SubrootGroup = () => {
    return (
        <div className="flex h-[75px] w-[200px] flex-col rounded-lg border-1.5 border-black bg-[#f0f0f040]">
            <Handle type="source" position={Position.Bottom} id="s-dev" style={{ left: '33%' }} />
            <Handle type="source" position={Position.Bottom} id="s-sec" style={{ left: '90%' }} />
            <Handle
                type="source"
                position={Position.Bottom}
                id="s-devops"
                style={{ left: '10%' }}
            />
            <Handle
                type="source"
                position={Position.Bottom}
                id="s-science"
                style={{ left: '66%' }}
            />
        </div>
    );
};
