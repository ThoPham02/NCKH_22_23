import client from "../apis";
import { useState, useEffect, useCallback } from "react";


function useApi(url) {
  const [data, setData] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchData = useCallback( async () => {
    try {
      setIsLoading(true);
      const response = await client.get(url);
      setData(response.data);
    } catch (error) {
      setError(error);
    } finally {
      setIsLoading(false);
    }
  }, [url]);

  useEffect(() => {
    fetchData()
  }, [fetchData]);

  return { data, isLoading, error, fetchData };
}

export default useApi;
