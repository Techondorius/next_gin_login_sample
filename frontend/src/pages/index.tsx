import axios from "axios"
import { url } from "inspector"
import { useEffect, useState } from "react"

const App = () => {
  interface indexResponse {
    message: string
  }
  const [res, setRes] = useState<indexResponse>()

  useEffect(() => {
    axios
      .get<indexResponse>("http://localhost:8080")
      .then((res) => {
        setRes(res.data)
      })
      .catch((e) => {})
  }, [])

  const getCookie = () => {
    axios
      .post(
        "http://localhost:8080/login",
        {
          UserID: "asdl0606",
          Password: "password",
        },
        {
          withCredentials: true,
        }
      )
      .catch((e) => {
        console.log(e)
      })
  }

  const checkCookie = () => {
    axios
      .get("http://localhost:8080/auth/aojiru", {
        withCredentials: true,
      })
      .then((res) => {
        console.log(res)
      })
  }

  return (
    <>
      <h1>{res ? res.message : "asdf"}</h1>
      <button onClick={getCookie}>GetCookie</button>
      <button onClick={checkCookie}>checkCookie</button>
    </>
  )
}

export default App
