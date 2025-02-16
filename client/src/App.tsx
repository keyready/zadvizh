import { Link, Route, Routes } from 'react-router';
import { useSelector } from 'react-redux';

import { FlowPage, MainPage, RatingPage, RegisterPage } from './pages';
import { RootState } from './app/store/store.ts';

function App() {
    const userAccessToken = useSelector((state: RootState) => state.user.accessToken);

    return (
        <div className="relative">
            {userAccessToken && (
                <div className="absolute left-0 right-0 top-0 z-50 flex w-full items-center justify-start gap-5 border-b-2 p-4">
                    <Link className="text-white" to="/hierarchy">
                        Структура Движа
                    </Link>
                    <Link className="text-white" to="/rating">
                        Рейтинг
                    </Link>
                </div>
            )}
            <Routes>
                <Route index element={<MainPage />} />
                {userAccessToken ? <Route path="/hierarchy" element={<FlowPage />} /> : null}
                {userAccessToken ? <Route path="/rating" element={<RatingPage />} /> : null}
                {/*<Route path="/hierarchy" element={<FlowPage />} />*/}
                {/*<Route path="/rating" element={<RatingPage />} />*/}
                <Route path="/auth/continue" element={<RegisterPage />} />
                <Route path="/*" element={<RegisterPage />} />
            </Routes>
        </div>
    );
}

export default App;
