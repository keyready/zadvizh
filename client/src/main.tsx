import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router';
import { HeroUIProvider } from '@heroui/system';
import { Provider } from 'react-redux';
import { ToastProvider } from '@heroui/react';

import { store } from './app/store/store';
import App from './App.tsx';

import './app/styles/index.scss';

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <Provider store={store}>
            <HeroUIProvider>
                <ToastProvider placement="top-center" />
                <BrowserRouter>
                    <App />
                </BrowserRouter>
            </HeroUIProvider>
        </Provider>
    </StrictMode>,
);
