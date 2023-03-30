import { useEffect } from "react";

function App() {
    useEffect (() => {
        fetch('/topics')
            .then((res) => res.json())
            .then((res) => console.log(res));
    },[])
}

export default App;