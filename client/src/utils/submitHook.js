import { useState, useEffect } from 'react'

function useSubmitFetch (url, opts) {
  const [response, setResponse] = useState(null)
  const [loading, setLoading] = useState(false)
  const [hasError, setHasError] = useState(false)
  const [hasRun, setHasRun] = useState(false)

  const execute = (data) => {
      console.log(data)
      //set default value
      !('gradient1' in data) && (data.gradient1 = "" )
      !('gradient2' in data) && (data.gradient2 = "")
      !('color' in data) && (data.color = "")

    setLoading(true)
    fetch(url, {
      method: 'post',
      body: JSON.stringify(data)
    })
      .then(res => {
        setResponse(res.data)
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
