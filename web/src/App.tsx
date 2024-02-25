import {useState} from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

// useCounterApi is a custom hook that will hit /api/increment, incrementing the server-side count and
// returning the new count.
function useCounterApi(initialState: number) {
    const [count, setCount] = useState(initialState)
    const increment = (by: number) => {
        fetch(`/api/increment/${by}`)
            .then((res) => res.json())
            .then(({count}) => setCount(count))
    }
    return {count, increment}

}

function App() {
    const {count, increment} = useCounterApi(0)

    return (
        <>
            <div>
                <a href="https://vitejs.dev" target="_blank">
                    <img src={viteLogo} className="logo" alt="Vite logo"/>
                </a>
                <a href="https://react.dev" target="_blank">
                    <img src={reactLogo} className="logo react" alt="React logo"/>
                </a>
            </div>
            <h1>Vite + React</h1>
            <div className="card">
                <button onClick={() => increment(7)}>
                    count is {count}
                </button>
                <p>
                    Edit <code>src/App.tsx</code> and save to test HMR
                </p>
            </div>
            <p className="read-the-docs">
                Click on the Vite and React logos to learn more
            </p>
        </>
    )
}

export default App
