// Copyright (c) 2014 Dataence, LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sequence

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	analyzerSshTests = []struct {
		msg, pat, patNoSp string
	}{
		{
			"Jan 12 06:49:42 irc sshd[7034]: Failed password for root from 218.161.81.238 port 4228 ssh2",
			"%regextime:1% %apphost% %appname% [ %sessionid% ] : %status% %method% for %srcuser% from %srcip% port %srcport% ssh2",
			"%regextime:1% %apphost% %appname%[%sessionid%]: %status% %method% for %srcuser% from %srcip% port %srcport% ssh2",
		},
		{
			"Jan 12 06:49:42 irc sshd[7034]: Accepted password for root from 218.161.81.238 port 4228 ssh2",
			"%regextime:1% %apphost% %appname% [ %sessionid% ] : %status% %method% for %srcuser% from %srcip% port %srcport% ssh2",
			"%regextime:1% %apphost% %appname%[%sessionid%]: %status% %method% for %srcuser% from %srcip% port %srcport% ssh2",
		},
		{
			"Jan 12 14:44:48 jlz sshd[11084]: Accepted publickey for jlz from 76.21.0.16 port 36609 ssh2",
			"%regextime:1% %apphost% %appname% [ %sessionid% ] : %status% %method% for %srcuser% from %srcip% port %srcport% ssh2",
			"%regextime:1% %apphost% %appname%[%sessionid%]: %status% %method% for %srcuser% from %srcip% port %srcport% ssh2",
		},
		{
			"Jan 12 08:03:01 buster-dev sshd[24877]: pam_unix(sshd:session): session opened for user jolata by (uid=0)",
			"%regextime:1% %apphost% %appname% [ %sessionid% ] : pam_unix ( sshd : %object% ) : session %action% for user %srcuser% by ( uid = %srcuid% )",
			"%regextime:1% %apphost% %appname%[%sessionid%]: pam_unix(sshd:%object%): session %action% for user %srcuser% by (uid=%srcuid%)",
		},
		{
			"Jan 12 11:05:23 irc sshd[7791]: pam_unix(sshd:auth): check pass; user unknown",
			"%regextime:1% %apphost% %appname% [ %sessionid% ] : pam_unix ( sshd : %string% ) : check pass ; user %srcuser%",
			"%regextime:1% %apphost% %appname%[%sessionid%]: pam_unix(sshd:%string%): check pass; user %srcuser%",
		},
	}

	analyzerKVTests = []struct {
		msg, pat, patNoSp string
	}{
		{
			"id=firewall time=\"2005-03-18 14:01:46\" fw=TOPSEC priv=6 recorder=kernel type=conn policy=414 proto=TCP rule=accept src=61.167.71.244 sport=35223 dst=210.82.119.211 dport=25 duration=27 inpkt=37 outpkt=39 sent=1770 rcvd=20926 smac=00:04:c1:8b:d8:82 dmac=00:0b:5f:b2:1d:80",
			"id = %string% time = \" %regextime:3% \" fw = %string% priv = %integer% recorder = %string% type = %string% policy = %integer% proto = %protocol% rule = %string% src = %srcip% sport = %srcport% dst = %dstip% dport = %dstport% duration = %duration% inpkt = %integer% outpkt = %integer% sent = %integer% rcvd = %integer% smac = %srcmac% dmac = %dstmac%",
			"id=%string% time=\"%regextime:3%\" fw=%string% priv=%integer% recorder=%string% type=%string% policy=%integer% proto=%protocol% rule=%string% src=%srcip% sport=%srcport% dst=%dstip% dport=%dstport% duration=%duration% inpkt=%integer% outpkt=%integer% sent=%integer% rcvd=%integer% smac=%srcmac% dmac=%dstmac%",
		},
		{
			"id=firewall time=\"2005-03-18 14:01:43\" fw=TOPSEC priv=4 recorder=kernel type=conn policy=504 proto=TCP rule=deny src=210.82.121.91 sport=4958 dst=61.229.37.85 dport=23124 smac=00:0b:5f:b2:1d:80 dmac=00:04:c1:8b:d8:82",
			"id = %string% time = \" %regextime:3% \" fw = %string% priv = %integer% recorder = %string% type = %string% policy = %integer% proto = %protocol% rule = %string% src = %srcip% sport = %srcport% dst = %dstip% dport = %dstport% smac = %srcmac% dmac = %dstmac%",
			"id=%string% time=\"%regextime:3%\" fw=%string% priv=%integer% recorder=%string% type=%string% policy=%integer% proto=%protocol% rule=%string% src=%srcip% sport=%srcport% dst=%dstip% dport=%dstport% smac=%srcmac% dmac=%dstmac%",
		},
	}

	analyzerSshdSamples []string = []string{
		"Jan 12 06:49:42 irc sshd[7034]: Failed password for root from 218.161.81.238 port 4228 ssh2",
		"Jan 12 06:49:42 irc sshd[7034]: Accepted password for root from 218.161.81.238 port 4228 ssh2",
		"Jan 12 14:44:48 jlz sshd[11084]: Accepted publickey for jlz from 76.21.0.16 port 36609 ssh2",
	}

	analyzerSshdPatterns []string = []string{
		"%time% %string% sshd [ %integer% ] : %string% %string% for %string% from %ipv4% port %integer% ssh2",
		"%time% %string% sshd [ %integer% ] : %string% %string% for %string% from %ipv4% port %integer% ssh2",
		"%time% %string% sshd [ %integer% ] : %string% %string% for %string% from %ipv4% port %integer% ssh2",
	}
)

func TestAnalyzerMergeNodes(t *testing.T) {
	atree := NewAnalyzer()
	scanner := NewScanner()
	var pos []int
	for _, data := range analyzerSshdSamples {
		seq, _, err := scanner.Scan(data, false, pos)
		require.NoError(t, err)
		err = atree.Add(seq)
		require.NoError(t, err)
	}

	l := 0

	for i := 1; i < allTypesCount; i++ {
		node := atree.levels[l][i]

		//added this to match the case the tag is known
		if node != nil && node.Tag != TagUnknown {
			require.NotNil(t, node, fmt.Sprintf("Expected: levels[%d][%d] != nil, Actual: got nil", l, i))
		} else {
			require.Nil(t, node, fmt.Sprintf("Expected: levels[%d][%d] == nil, Actual: got non-nil %s", l, i, node))
		}
	}

	require.Equal(t, 0, len(atree.litmaps[l]), fmt.Sprintf("Expected: levels[%d].litmap == 0, Actual: got %d", l, len(atree.litmaps[l])))

	l = 1

	for i := 1; i < allTypesCount; i++ {
		node := atree.levels[l][i]

		require.Nil(t, node, fmt.Sprintf("Expected: levels[%d][%d] == nil, Actual: got non-nil %s", l, i, node))
	}

	if len(atree.litmaps[l]) != 2 {
		t.Fatalf("Expected: levels[%d].litmap == 2, Actual: got %d", l, len(atree.litmaps[l]))
	}

	l = 2

	for i := 1; i < allTypesCount; i++ {
		node := atree.levels[l][i]

		require.Nil(t, node, fmt.Sprintf("Expected: levels[%d][%d] == nil, Actual: got non-nil %s", l, i, node))
	}

	require.Equal(t, 1, len(atree.litmaps[l]), fmt.Sprintf("Expected: levels[%d].litmap == 1, Actual: got %d", l, len(atree.litmaps[l])))

	l = 4

	for i := 1; i < allTypesCount; i++ {
		node := atree.levels[l][i]

		if i == TagTypesCount+int(TokenInteger) {
			require.NotNil(t, node, fmt.Sprintf("Expected: levels[%d][TokenInteger] != nil, Actual: got nil", l))
		} else {
			require.Nil(t, node, fmt.Sprintf("Expected: levels[%d][%d] == nil, Actual: got non-nil %s", l, i, node))
		}
	}

	require.Equal(t, 0, len(atree.litmaps[l]), fmt.Sprintf("Expected: levels[%d].litmap == 0, Actual: got %d", l, len(atree.litmaps[l])))

	l = 7

	for i := 1; i < allTypesCount; i++ {
		node := atree.levels[l][i]

		require.Nil(t, node, fmt.Sprintf("Expected: levels[%d][%d] == nil, Actual: got non-nil %s", l, i, node))
	}

	require.Equal(t, 2, len(atree.litmaps[l]), fmt.Sprintf("Expected: levels[%d].litmap == 2, Actual: got %d", l, len(atree.litmaps[l])))

	l = 8

	for i := 1; i < allTypesCount; i++ {
		node := atree.levels[l][i]

		require.Nil(t, node, fmt.Sprintf("Expected: levels[%d][%d] == nil, Actual: got non-nil %s", l, i, node))
	}

	require.Equal(t, 2, len(atree.litmaps[l]), fmt.Sprintf("Expected: levels[%d].litmap == 2, Actual: got %d", l, len(atree.litmaps[l])))

	l = 12

	for i := 1; i < allTypesCount; i++ {
		node := atree.levels[l][i]

		if i == TagTypesCount+int(TokenIPv4) {
			require.NotNil(t, node, fmt.Sprintf("Expected: levels[%d][TokenIPv4] != nil, Actual: got nil", l))
		} else {
			require.Nil(t, node, fmt.Sprintf("Expected: levels[%d][%d] == nil, Actual: got non-nil %s", l, i, node))
		}
	}

	require.Equal(t, 0, len(atree.litmaps[l]), fmt.Sprintf("Expected: levels[%d].litmap == 0, Actual: got %d", l, len(atree.litmaps[l])))

	l = 14

	for i := 1; i < allTypesCount; i++ {
		node := atree.levels[l][i]

		if i == TagTypesCount+int(TokenInteger) {
			require.NotNil(t, node, fmt.Sprintf("Expected: levels[%d][TokenInteger] != nil, Actual: got nil", l))
		} else {
			require.Nil(t, node, fmt.Sprintf("Expected: levels[%d][%d] == nil, Actual: got non-nil %s", l, i, node))
		}
	}

	require.Equal(t, 0, len(atree.litmaps[l]), fmt.Sprintf("Expected: levels[%d].litmap == 0, Actual: got %d", l, len(atree.litmaps[l])))

	atree.Finalize()

	for _, l := range []int{1, 7, 8, 10} {
		require.Equal(t, allTypesCount+1, len(atree.levels[l]), fmt.Sprintf("Expected: len(levels[%d]) == %d, Actual: got non-nil %d", l, allTypesCount+1, len(atree.levels[l])))

		for i := 1; i < allTypesCount; i++ {
			node := atree.levels[l][i]
			require.Nil(t, node, fmt.Sprintf("Expected: levels[%d][%d] == nil, Actual: got non-nil %s", l, i, node))
		}

		node := atree.levels[l][allTypesCount]
		require.Equal(t, TokenString, node.Type, fmt.Sprintf("Expected: levels[%d][%d].Type == TokenString, Actual: got %s", l, allTypesCount+1, node.Type))
	}
}

func TestAnalyzerMatchPatterns(t *testing.T) {
	atree := NewAnalyzer()
	scanner := NewScanner()
	var pos []int

	for _, tc := range analyzerSshTests {
		seq, _, err := scanner.Scan(tc.msg, false, pos)
		require.NoError(t, err)
		err = atree.Add(seq)
		require.NoError(t, err, tc.msg)
	}

	for _, tc := range analyzerKVTests {
		seq, _, err := scanner.Scan(tc.msg, false, pos)
		require.NoError(t, err)
		err = atree.Add(seq)
		require.NoError(t, err, tc.msg)
	}

	atree.Finalize()

	for _, tc := range analyzerSshTests {
		seq, _, err := scanner.Scan(tc.msg, false, pos)
		require.NoError(t, err)
		seq, err = atree.Analyze(seq)
		require.NoError(t, err, tc.msg)
		r, _ := seq.String()
		if config.markSpaces {
			require.Equal(t, tc.patNoSp, r, tc.msg+"\n"+seq.PrintTokens())
		} else {
			require.Equal(t, tc.pat, r, tc.msg+"\n"+seq.PrintTokens())
		}

	}

	for _, tc := range analyzerKVTests {
		seq, _, err := scanner.Scan(tc.msg, false, pos)
		require.NoError(t, err)
		seq, err = atree.Analyze(seq)
		require.NoError(t, err, tc.msg)
		r, _ := seq.String()
		if config.markSpaces {
			require.Equal(t, tc.patNoSp, r, tc.msg+"\n"+seq.PrintTokens())
		} else {
			require.Equal(t, tc.pat, r, tc.msg+"\n"+seq.PrintTokens())
		}
	}
}
