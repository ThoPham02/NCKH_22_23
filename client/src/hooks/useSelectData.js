import { useEffect } from "react"
import { useDispatch, useSelector } from "react-redux"

export const useSelectData = ({action, selector, payload}) => {
    const data = useSelector(selector)

    const dispatch = useDispatch()

    useEffect(() => {
        if (data) {
            dispatch(action(payload))
        }
    }, [dispatch, payload, data])

    return data
}