import {useNavigate} from "react-router-dom"

function Login(){
    const navigate = useNavigate()
    
    function secondpage(){
        navigate('/second')
    }

    return (
        <div className="space-y-4 p-6 bg-slate-300 rounded-md shadow flex flex-col">
        <input type="text" placeholder="Usuário" className="border border-slate-300 outline-slate-950 rounded-md"></input>
        <button onClick={(secondpage)} className="bg-slate-700 text-white px-4 py-2 rounded-md font-medium">Logar com Github</button>

    </div>
    );
}

export default Login