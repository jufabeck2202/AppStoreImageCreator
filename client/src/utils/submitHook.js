import { useState, useEffect } from 'react'

function useSubmitFetch (url) {
  const [response, setResponse] = useState(null)
  const [loading, setLoading] = useState(false)
  const [hasError, setHasError] = useState(false)
  const [hasRun, setHasRun] = useState(false)

  const execute = data => {
    //set default value
    !('gradient1' in data) && (data.gradient1 = '')
    !('gradient2' in data) && (data.gradient2 = '')
    !('color' in data) && (data.color = '')
    console.log(data)

    setLoading(true)
    fetch(url, {
      method: 'post',
      headers: {
        'Accept': 'application/json, text/plain, */*',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    }).then(res=>res.json())
      .then(res => {
        setResponse(res)
        console.log(res)
        setLoading(false)
        setHasRun(true)
      })
      .catch(() => {
        setHasError(true)
        setLoading(false)
      })
  }
  return [execute, response, loading, hasError]
}

export default useSubmitFetch
