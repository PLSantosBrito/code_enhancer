import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App.jsx'
import './index.css'
import {createBrowserRouter, RouterProvider} from 'react-router-dom';
import PageLoged from './pages/PageLoged.jsx';

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />
  },
  {
    path: '/home',
    element: <PageLoged/>
  }
]);


createRoot(document.getElementById('root')).render(
  
    <RouterProvider router={router}/>
  
)
