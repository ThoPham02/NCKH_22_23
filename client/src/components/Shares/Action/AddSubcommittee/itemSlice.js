import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../../apis";

const ItemSlice = createSlice({
    name: "ItemSlice",
    initialState: {
        item: [
            {
                index: 0,
                lectureID: 0,
                lectureName: "",
                role: 1,
            }
        ],
        status: "idle"
    },
    reducers: {
        changeItem: (state, action) => {
            const { index, lectureID, lectureName, role } = action.payload;
            const itemIndex = state.item.findIndex(item => item.index === index);
            if (itemIndex !== -1) {
                state.item[itemIndex] = {
                    index,
                    lectureID,
                    lectureName,
                    role,
                };
            }
        },
        addItem: (state, action) => {
            const newItem = {index: state.item.length, lectureID: 0, lectureName: "", role: 1}
            state.item.push(newItem);
        },
        resetItem: (state, action) => {
            state.item = [
                {
                    index: 0,
                    lectureID: 0,
                    lectureName: "",
                    role: 1,
                }
            ]
        }
    },
    extraReducers: builder => {
        builder
        .addCase(fetchTeacher.pending, (state, action) => {
            state.status = "loading"
        })
        .addCase(fetchTeacher.fulfilled, (state, action) => {
            state.status = "idle"
            state.teacher = action.payload.members
        })
    }
})

export const fetchTeacher = createAsyncThunk("fetchTeacher", async ({facultyID}) => {
    const resp = await client.get("/user/subcommittee", {
        params: {
            role: 2,
            faculty_id: facultyID,
        }
    })
    return resp.data
})

export default ItemSlice;
export const ItemReducer =  ItemSlice.reducer;
export const ItemAction =  ItemSlice.actions;