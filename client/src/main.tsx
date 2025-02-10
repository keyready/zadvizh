import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import App from './App.tsx';
import { BrowserRouter } from 'react-router';
import { HeroUIProvider } from '@heroui/system';

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
