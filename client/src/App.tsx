import { Route, Routes } from 'react-router';

import { FlowPage, MainPage, RegisterPage } from './pages';

function App() {
    return (
        <Routes>
            <Route index element={<MainPage />} />
            <Route path="/auth/continue" element={<RegisterPage />} />
            <Route path="/hierarchy" element={<FlowPage />} />
            <Route path="/*" element={<RegisterPage />} />
        </Routes>
    );
}

export default App;
