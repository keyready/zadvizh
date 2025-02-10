import { Routes, Route } from 'react-router';
import { MainPage, RegisterPage } from './pages';

function App() {
    return (
        <Routes>
            <Route index element={<MainPage />} />
            <Route path="auth/continue" element={<RegisterPage />} />
        </Routes>
    );
}

export default App;
