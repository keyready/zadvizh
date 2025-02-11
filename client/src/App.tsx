import { Route, Routes } from 'react-router';

import { FlowPage, MainPage, RegisterPage } from './pages';

function App() {
    // const token = localStorage.getItem('t');

    return (
        <Routes>
            <Route index element={<MainPage />} />
            {/*{token ? <Route path="/hierarchy" element={<FlowPage />} /> : null}*/}
            <Route path="/hierarchy" element={<FlowPage />} />
            <Route path="/auth/continue" element={<RegisterPage />} />
            <Route path="/*" element={<RegisterPage />} />
        </Routes>
    );
}

export default App;
