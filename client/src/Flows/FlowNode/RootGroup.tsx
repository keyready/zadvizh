import { Handle, Position } from '@xyflow/react';
import '@xyflow/react/dist/style.css';

export const RootGroup = () => {
    return (
        <div className="flex h-[150px] w-[400px] flex-col rounded-lg border-1.5 border-black bg-[#f0f0f040]">
            <Handle type="source" position={Position.Bottom} id="dev" style={{ left: '33%' }} />
            <Handle type="source" position={Position.Bottom} id="sec" style={{ left: '90%' }} />
            <Handle type="source" position={Position.Bottom} id="devops" style={{ left: '10%' }} />
            <Handle type="source" position={Position.Bottom} id="science" style={{ left: '66%' }} />
        </div>
    );
};
