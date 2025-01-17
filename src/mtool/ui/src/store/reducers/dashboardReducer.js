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

import isEqual from 'lodash/isEqual';
import * as actionTypes from "../actions/actionTypes"

const initialState = {
    ip: '0.0.0.0',
    mac: 'NA',
    volumes: [],
    arrayVolumes: [],
    selectedArray: 'all',
    arraySize: 0,
    arrayVols: {},
    showTelemetryAlert: false,
    errorMsg: "",
    readIOPS: 0,
    writeIOPS: 0,
    readBW: 0,
    writeBW: 0,
    readLatency: 0,
    writeLatency: 0,
    devices: [],
    errorInDevices: true,
    totalNominals: 0,
    totalWarnings: 0,
    totalCriticals: 0,
}


const dashboardReducer = (state = initialState, action) => {
    switch (action.type) {
        case actionTypes.CLOSE_TELEMETRY_ALERT:
            return {
                ...state,
                showTelemetryAlert: false
            }
        case actionTypes.SET_SHOW_TELEMETRY_NOT_RUNNING:
            return {
                ...state,
                showTelemetryAlert: action.showTelemetryAlert,
                errorMsg: action.errorMsg
            };
        case actionTypes.SELECT_ARRAY: {
            let arrayVolumes = [];
            if (action.array === "all") {
                arrayVolumes = [...state.volumes]
            } else {
                state.volumes.forEach((volume) => {
                    if (volume.array === action.array) {
                        arrayVolumes.push({
                            ...volume
                        });
                    }
                });
            }
            return {
                ...state,
                selectedArray: action.array,
                arrayVolumes
            }
        }
        case actionTypes.FETCH_VOLUME_INFO: {
            const volumes = [];
            const arrayVols = {};
            let arrayVolumes = [];
            Object.keys(action.volumes).forEach((array) => {
                arrayVols[array] = action.volumes[array].length;
                action.volumes[array].forEach((vol) => {
                    volumes.push({
                        ...vol,
                        array
                    });
                    if (state.selectedArray !== "all" && array === state.selectedArray) {
                        arrayVolumes.push({
                            ...volumes[volumes.length - 1]
                        })
                    }
                });
            });
            if (state.selectedArray === "all") {
                arrayVolumes = [...volumes];
            }
            return {
                ...state,
                volumes,
                arrayVols,
                arrayVolumes
            };
        }
        case actionTypes.FETCH_PERFORMANCE_INFO:
            return {
                ...state,
                readIOPS: action.readIOPS,
                writeIOPS: action.writeIOPS,
                readBW: action.readBW,
                writeBW: action.writeBW,
                readLatency: action.readLatency,
                writeLatency: action.writeLatency
            };
        case actionTypes.FETCH_HARDWARE_HEALTH: {
            const newState = {
                ...state,
                devices: action.devices,
                errorInDevices: action.errorInDevices,
                totalNominals: action.totalNominals,
                totalWarnings: action.totalWarnings,
                totalCriticals: action.totalCriticals,
            }
            if (state.devices.length !== action.devices.length)
                return newState
            for (let i = 0; i < state.devices.length; i += 1) {
                if (!isEqual(state.devices[i], action.devices[i]))
                    return newState
            }
            if (
                action.errorInDevices !== state.errorInDevices ||
                action.totalCriticals !== state.totalCriticals ||
                action.totalWarnings !== state.totalWarnings ||
                action.totalNominals !== state.totalNominals 
            )
                return newState;
            return state;
        }
        case actionTypes.FETCH_IPANDMAC_INFO:
            return {
                ...state,
                ip: action.ip,
                mac: action.mac
            };
        default:
            return state;
    }
};

export default dashboardReducer;
