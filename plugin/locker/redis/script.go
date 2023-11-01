package redis

var AcquireLockScript = `
local ret
local remote_id = KEYS[1]
local remote_signature = ARGV[1]

if type(remote_id) ~= 'string' then 
    return {'0', 'Bad id'}
end
if string.len(remote_id) <= 0 then
   return {'0', 'Bad id'}
end
if type(remote_signature) ~= 'string' then 
    return {'0', 'Bad Signature'}
end
if string.len(remote_signature) ~= 64 then
   return {'0', 'Bad Signature'}
end

ret = redis.call('setnx', remote_id, remote_signature)
if ret == 1 then
   return {'1', 'success'}
end

local native_signature = redis.call('get', remote_id)
local native_signature_len = string.len(native_signature)
if native_signature_len ~= 64 then
   return {'0', 'The previous signature is incorrect'}
end

local temp = string.sub(native_signature, 52, 65)
local native_to = tonumber(temp)
if type(native_to) ~= 'number' then
   return {'0', 'The type of to is not number'}
end

local temp = string.sub(remote_signature, 38, 50)
local remote_from = tonumber(temp)
if native_to > remote_from then
   return {'0', 'The previous lock has not expired yet!', native_to, remote_from}
end

redis.call('set', remote_id, remote_signature)
return {'1', 'The previous lock has expired!'}
`

var ReleaseLockScript = `
local remote_id = KEYS[1]
local remote_signature = ARGV[1]

if type(remote_id) ~= 'string' then 
    return {'0', 'Bad id'}
end
if string.len(remote_id) <= 0 then
   return {'0', 'Bad id'}
end
if type(remote_signature) ~= 'string' then 
    return {'0', 'Bad Signature'}
end
if string.len(remote_signature) ~= 64 then
   return {'0', 'Bad Signature'}
end

local native_signature = redis.call('get', remote_id)
local native_signature_len = string.len(native_signature)
if native_signature_len ~= 64 then
   return {'0', 'The previous signature is incorrect', remote_id, remote_signature}
end

local native_uuid = string.sub(native_signature, 1, 36)
local remote_uuid = string.sub(remote_signature, 1, 36)
if native_uuid == remote_uuid then
   redis.call('del', remote_id)
   return {'1', 'success'}
end

return {'0', 'release fail, operation not permitted'}
`

var RefreshLockScript = `
local remote_id = KEYS[1]
local remote_signature = ARGV[1]

if type(remote_id) ~= 'string' then 
    return {'0', 'Bad id'}
end
if string.len(remote_id) <= 0 then
   return {'0', 'Bad id'}
end
if type(remote_signature) ~= 'string' then 
    return {'0', 'Bad Signature'}
end
if string.len(remote_signature) ~= 64 then
   return {'0', 'Bad Signature'}
end

local native_signature = redis.call('get', remote_id)
local native_signature_len = string.len(native_signature)
if native_signature_len ~= 64 then
   return {'0', 'The previous signature is incorrect'}
end

local native_uuid = string.sub(native_signature, 1, 36)
local remote_uuid = string.sub(remote_signature, 1, 36)
if native_uuid == remote_uuid then
   redis.call('set', remote_id, remote_signature)
   return {'1', 'success'}
end

return {'0', 'refresh fail, operation not permitted'}
`
