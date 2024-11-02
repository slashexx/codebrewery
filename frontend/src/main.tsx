import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'; // Make sure to import BrowserRouter
import App from './App.tsx';
import './index.css';
import { IntroPage } from './IntroPage.tsx';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Router> {/* Wrap Routes with Router */}
      <Routes>
        <Route path="/" element={<IntroPage />} />
        <Route path="/editor" element={<App />} />
      </Routes>
    </Router>
  </StrictMode>
);
