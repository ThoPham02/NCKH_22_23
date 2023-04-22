import { configureStore } from '@reduxjs/toolkit'
import { LoginReducer } from '../pages/common/Login/loginSlice'

const store = configureStore({
  reducer: {
    login: LoginReducer
  }
})

export default store