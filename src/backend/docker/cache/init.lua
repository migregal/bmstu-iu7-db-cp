box.cfg{listen = os.getenv('CACHE_PORT')}

box.once("bootstrap", function()
    box.schema.user.create(os.getenv('CACHE_DATABASE_USERNAME'), { password = os.getenv('CACHE_DATABASE_USER_PASSWORD'), if_not_exists = true })
    box.schema.user.grant(os.getenv('CACHE_DATABASE_USERNAME'), 'read,write,execute,session,usage,create,drop,alter,reference,trigger,insert,update,delete', 'universe', nil, { if_not_exists = true })

    local auth_cache_space = box.schema.space.create(
        os.getenv('CACHE_DATABASE_MODEL_SPACE'),
        { if_not_exists = true }
    )

    auth_cache_space:format({
        { name = 'key',   type = 'string' },
        { name = 'value', type = '*' },
    })

    auth_cache_space:create_index('primary',
        { type = 'hash', parts = {1, 'string'}, if_not_exists = true }
    )
end)
