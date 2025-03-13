import React, {useEffect, useState} from "react";


function PageLoged(){
    

    return (
        <div className="h-screen flex flex-col bg-slate-800">
    
          <header className="bg-slate-900 text-white p-6 text-xl font-bold text-center shadow-xl">
            Header
          </header>
    
          <div className="flex flex-1">
            <aside className="w-60 bg-slate-900 p-5 text-white h-screen fixed lg:relative border-r border-slate-700">
              <h2 className="text-xl font-bold">Projetos</h2>
              <ul className="mt-4 space-y-2">
                <li className="hover:bg-slate-700 p-2 rounded cursor-pointer transition-all duration-300 ease-in-out">
                  Projeto 1
                </li>
                <li className="hover:bg-slate-700 p-2 rounded cursor-pointer transition-all duration-300 ease-in-out">
                  Projeto 2
                </li>
                <li className="hover:bg-slate-700 p-2 rounded cursor-pointer transition-all duration-300 ease-in-out">
                  Projeto 3
                </li>
              </ul>
            </aside>
    
        <main className="flex-1 p-10 text-white ml-72 lg:ml-0 lg:pl-12 overflow-auto">

          <div className="bg-slate-700 h-full w-full rounded-lg p-10 shadow-lg">
            <h2 className="text-2xl font-semibold"> Código aqui</h2>
          </div>
        </main>
          </div>
        </div>
      );
}

export default PageLoged;