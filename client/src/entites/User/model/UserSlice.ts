import type { PayloadAction } from '@reduxjs/toolkit';
import { createSlice } from '@reduxjs/toolkit';

export interface UserState {
    id: string;
    accessToken?: string;
}

const initialState: UserState = {
    id: '',
    accessToken: '',
};

export const UserSlice = createSlice({
    name: 'counter',
    initialState,
    reducers: {
        setUserTgId: (state, action: PayloadAction<string>) => {
            state.id = action.payload;
        },
        setUserAccessToken: (state, action: PayloadAction<string>) => {
            state.accessToken = action.payload;
        },
    },
});

export const { setUserTgId, setUserAccessToken } = UserSlice.actions;
export const UserReducer = UserSlice.reducer;
