// import { configureStore } from '@reduxjs/toolkit';
// import rootReducer from './reducer';
// import { composeWithDevTools } from 'redux-devtools-extension';

// const composedEnhancers = composeWithDevTools({});

// const store = configureStore({
//   reducer:
//   rootReducer,
//   composedEnhancers
// });

// export default store;

import { configureStore } from '@reduxjs/toolkit'
import todoSlice from '../pages/Topic/components/TodoList/todoSlice.js'
import filtersSlice from '../pages/Topic/components/Filters/filtersSlice.js'

const store = configureStore({
  reducer: {
    filters: filtersSlice.reducer,
    todoList: todoSlice.reducer
  }
})

export default store