import { configureStore } from '@reduxjs/toolkit'
import { loginReducer } from '../pages/Login/LoginSlice.js'
import { FilterReducer } from '../components/Shares/Search/SearchSlice.js'

const store = configureStore({
  reducer: {
    topicFilter: FilterReducer,
    login: loginReducer
  }
})

export default store