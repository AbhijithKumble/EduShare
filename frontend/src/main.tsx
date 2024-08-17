import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import Routing from './router/Router.tsx'


import App from './App.tsx'
import './index.css'


ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <BrowserRouter>
           <App> 
                <Routing />
            </App>
        </BrowserRouter>
    </React.StrictMode>,
);
