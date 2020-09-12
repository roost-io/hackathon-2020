set -x
sudo apt-get install -y jq curl
echo 'Before vote'
curl roost-controlplane:30080 | jq
echo 'Vote now'
curl -d '{
       "candidate_id": "roost",
       "voter_id": "unique_id_123"
}' -H 'Content-Type: application/json' roost-controlplane:30080 | jq
echo 'After Voting'
curl roost-controlplane:30080 | jq
