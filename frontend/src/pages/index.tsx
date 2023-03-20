import axios from "axios"
import { useEffect, useState } from "react"

const App = () => {
  interface indexResponse {
    message: string
  }
  const [res, setRes] = useState<indexResponse>()

  useEffect(() => {
    axios.get<indexResponse>("http://localhost:8080").then((res) => {
      setRes(res.data)
    }).catch((e) => {
      
    })
  }, [])

  return <h1>{res ? res.message : "asdf"}</h1>
}

export default App
