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

export { 
    updateTimestamp,
    asyncIsiBOFOSRunning,
    setOperationsMessage,
    setPOSInfo,
    setPOSProperty,
    setIsStatusCheckDone
} from "./headerActions";

export { 
    setShowTelemetryNotRunning,
    enableFetchingAlerts,
    fetchVolumes, 
    fetchAlerts,
    fetchPerformance,
    fetchStorage,
    fetchIpAndMac,
} from "./dashboardActions";
export {
    openAlertBox,
    // setAlertsInfo,
    fetchAlertsInfo,
    fetchAlertsType,
} from "./alertManagementActions";

// export { 
//     changeLanguage
//  } from "./headerLanguageActions";

export { 
    changeEmailList,
    fetchEmailList,
    setAlertBox,
    changeSmtpServer,
    setSmtpServer,
    deleteConfiguredSmtpServer,
    setIbofOSTimeInterval,
 } from "./configurationsettingActions";

 export { 
    getIbofOsLogs,
    setLiveLogs,
 } from "./logManagementActions";

export {
    fetchDevices,
    fetchDeviceDetails,
    fetchArray,
    fetchArrayDetails,
    fetchConfig,
    fetchArraySize,
    setNoArray,
    showStorageAlert,
    fetchStorageVolumes,
    startStorageLoader,
    stopStorageLoader,
    fetchMaxVolumeCount,
    toggleCreateVolumeButton,
    addVolumeDetails,
    clearVolumes,
    startFetchingVolumes,
    stopFetchingVolumes
} from "./storageActions";

export {
    toggleAdvanceCreateVolumePopup,
    resetInputs,
} from "./createVolumeActions";

export {
    getSubsystems,
    showSubsystemAlert
} from "./subsystemActions";

export {
    // fetchDiskUsed,
    // fetchDiskWrite,
    fetchCpuUsage,
    fetchReadBandwidth,
    fetchWriteBandwidth,
    fetchReadIops,
    fetchWriteIops,
    fetchVolReadBandwidth,
    fetchVolWriteBandwidth,
    fetchVolReadIops,
    fetchVolWriteIops,
    fetchReadLatency,
    fetchVolReadLatency,
    fetchWriteLatency,
    fetchVolWriteLatency,
    // fetchInputPowerVariation
} from "./performanceActions";

export {
    setIsResettingConfig,
    setIsResetConfigFailed,
    setIsSavingConfig,
    setShowConfig,
    setIsConfigured,
    setconfigurationFailed,
    changeCredentials,
    setIsLoggedIn,
    setLoginFailed,
    resetIsLoggedIn,
} from "./authenticationActions";

export {
    // setUsersInfo,
    fetchUsersInfo,
} from "./userManagementActions";

export {
    fetchServerInfo,
    fetchPowerInfo,
    fetchChassisFrontInfo,
    // fetchChassisRearInfo,
} from "./hardwareOverviewActions";

export {
    fetchPowerSensorInfo,
    fetchFanSensorInfo,
    fetchTemperatureSensorInfo,
} from "./hardwareSensorActions";

export {
    fetchSoftwareDetails,
    fetchHardwareDetails,
    fetchNetworkDetails,
} from "./hardwareHealthActions";

export {
    fetchPowerSummary,
} from "./hardwarePowerManagementActions";

export {
    startLoader,
    stopLoader,
} from "./waitLoaderActions";

export {
    BMCChangeCredentials,
    // BMCResetIsLoggedIn,
    BMCSetIsLoggedIn,
    BMCSetLoginFailed,
} from "./BMCAuthenticationActions";

export {
    setTelemetryStatus,
    setTelemetryProperties,
    openTelemetryAlert,
} from "./telemetryActions";