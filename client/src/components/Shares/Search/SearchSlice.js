import { createSlice } from "@reduxjs/toolkit";

const FilterSlice =  createSlice({
    name: 'filters',
    initialState: {
        search: '',
        dateFrom: '',
        dateTo: '',
        status: 0,
        department: 0,
        faculity: 0
    },
    reducers: {
        searchFilterChange(state, action) {
            state = action.payload;
        },
    }
})

export default FilterSlice;

export const FilterReducer =  FilterSlice.reducer;