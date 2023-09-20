import './App.css';
import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import Chat from './Pages/Chat';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={<Chat />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
