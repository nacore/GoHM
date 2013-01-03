/* The copyright in this software is being made available under the BSD
 * License, included below. This software may be subject to other third party
 * and contributor rights, including patent rights, and no such rights are
 * granted under this license.
 *
 * Copyright (c) 2012-2013, H265.net
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *  * Redistributions of source code must retain the above copyright notice,
 *    this list of conditions and the following disclaimer.
 *  * Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *  * Neither the name of the H265.net nor the names of its contributors may
 *    be used to endorse or promote products derived from this software without
 *    specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS
 * BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF
 * THE POSSIBILITY OF SUCH DAMAGE.
 */

package main

import (
    "fmt"
    "gohm/TAppDecoder"
    "gohm/TLibCommon"
    "log"
    "os"
    "time"
)

func main() {
    fmt.Printf("GoHM Software: Decoder Version [%s]\n", TLibCommon.NV_VERSION)

    cTAppDecTop := TAppDecoder.NewTAppDecTop()

    // create application decoder class
    cTAppDecTop.Create()

    if len(os.Args) == 1 {
        var args = []string{"gohm.exe", "test.bin", "test.yuv", "100", "trace.txt"}

        // parse configuration
        if err := cTAppDecTop.ParseCfg(len(args), args); err != nil {
            log.Fatal(err)
            return
        }
    } else {
        // parse configuration
        if err := cTAppDecTop.ParseCfg(len(os.Args), os.Args); err != nil {
            log.Fatal(err)
            return
        }
    }

    // starting time
    lBefore := time.Now()

    // call decoding function
    cTAppDecTop.Decode()

    // ending time
    lAfter := time.Now()

    fmt.Printf("\n\nTotal Decoding Time: %v.\n", lAfter.Sub(lBefore))

    // destroy application decoder class
    cTAppDecTop.Destroy()
}
