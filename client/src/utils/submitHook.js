import { useState, useEffect } from 'react'

function useSubmitFetch(url, opts) {
    const [response, setResponse] = useState(null)
    const [loading, setLoading] = useState(false)
    const [hasError, setHasError] = useState(false)
    const [hasRun, setHasRun] = useState(false)

    const execute = () => {    
    setLoading(true)
    fetch(url, opts)
        .then((res) => {
        setResponse(res.data)
        setLoading(false)
        setHasRun(true)
    })
        .catch(() => {
            setHasError(true)
            setLoading(false)
        })
    }
    return [execute, response, loading, hasError ]
}

export default useSubmitFetch