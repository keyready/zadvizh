import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router';
import { HeroUIProvider } from '@heroui/system';

import App from './App.tsx';

import './app/styles/index.scss';

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <HeroUIProvider>
            <BrowserRouter>
                <App />
            </BrowserRouter>
        </HeroUIProvider>
    </StrictMode>,
);
