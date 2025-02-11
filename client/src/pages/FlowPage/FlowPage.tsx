import { ReactFlowProvider } from '@xyflow/react';

import { FlowTree } from '../../Flows/FlowTree/FlowTree.tsx';

export const FlowPage = () => {
    return (
        <section>
            <ReactFlowProvider>
                <FlowTree />
            </ReactFlowProvider>
        </section>
    );
};
