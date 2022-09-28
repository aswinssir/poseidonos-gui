/*
 *   BSD LICENSE
 *   Copyright (c) 2021 Samsung Electronics Corporation
 *   All rights reserved.
 *
 *   Redistribution and use in source and binary forms, with or without
 *   modification, are permitted provided that the following conditions
 *   are met:
 *
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in
 *       the documentation and/or other materials provided with the
 *       distribution.
 *     * Neither the name of Samsung Electronics Corporation nor the names of its
 *       contributors may be used to endorse or promote products derived
 *       from this software without specific prior written permission.
 *
 *   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 *   "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 *   LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 *   A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 *   OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 *   SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 *   LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 *   DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 *   THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 *   (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 *   OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

import * as actionTypes from "../actions/actionTypes"

export const initialState = {
    timestamp:"...",
    status:false,
    state: "",
    OS_Running_Status: "...",
    operationsMessage: [],
    posVersion: "...",
    posProperty: "...",
    isStatusCheckDone: true
}


const headerReducer = (state = initialState, action) => {
    switch (action.type) {
        case actionTypes.GET_IS_IBOF_OS_RUNNING:
            return {
                ...state,
                status: action.status,
                OS_Running_Status: action.OS_Running_Status,
                state: action.state
            };
        case actionTypes.UPDATE_TIMESTAMP:
            return {
                ...state,
                timestamp: action.timestamp,
            };
        case actionTypes.SET_OPERATIONS_MESSAGE:
            return {
                ...state,
                operationsMessage: action.message
            }
        case actionTypes.SET_IS_STATUS_CHECK_DONE:
            return {
                ...state,
                isStatusCheckDone: action.payload
            }
	case actionTypes.SET_POS_INFO:
	    return {
		...state,
		posVersion: action.payload ? action.payload.version : ""
	    }
        case actionTypes.SET_POS_PROPERTY:
            return {
                ...state,
                posProperty: action.payload.rebuildPolicy
            }
        default:
            return state;
    }
};

export default headerReducer;
