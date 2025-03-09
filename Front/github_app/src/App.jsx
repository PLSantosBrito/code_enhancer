import { useCallback, useEffect, useState, useMemo } from "react";
import Login from "./components/Login"

function App(){

  return(
    <div className="w-screen h-screen bg-slate-700 flex justify-center p-6 ">
      <div className="w-[500]">
      <h1 className = "text-slate-100 text-3xl"> Github</h1>
      <Login/>
      </div>
    </div>
  );

}

export default App;