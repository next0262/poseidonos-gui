'''
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
 *     * Neither the name of Intel Corporation nor the names of its
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
 '''
 
from logging import StreamHandler
from util.db.influx import get_connection
import datetime

DATABASE = 'poseidon'
IP_INDEX = 1
METHOD_INDEX = 2
SCHEME_INDEX = 3
PATH_INDEX = 4
STATUS_INDEX = 5

class InfluxHandler(StreamHandler):

    def __init__(self):
        StreamHandler.__init__(self)
        self.influx_client = get_connection()

    def emit(self, record):
        fields = {
            "value": record.getMessage(),
            "path": record.args[PATH_INDEX],
            "status": record.args[STATUS_INDEX]
        }
        tags = {
            "level": record.levelname,
            "entity": "MTOOL",
            "ip": record.args[IP_INDEX],
            "method": record.args[METHOD_INDEX],
            "scheme": record.args[SCHEME_INDEX]
        }
        data = [{
            "fields": fields,
            "tags": tags,
            "measurement": "mtool_log",
            "time": datetime.datetime.now().isoformat()
        }]
        self.influx_client.write_points(points=data, database="poseidon", retention_policy="autogen", protocol="json")
