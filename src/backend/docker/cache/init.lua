box.cfg{listen = os.getenv('CACHE_PORT')}
box.cfg{memtx_max_tuple_size=64*1048576}

box.once("bootstrap", function()
    box.schema.user.create(os.getenv('CACHE_DATABASE_USERNAME'), { password = os.getenv('CACHE_DATABASE_USER_PASSWORD'), if_not_exists = true })
    box.schema.user.grant(os.getenv('CACHE_DATABASE_USERNAME'), 'read,write,execute,session,usage,create,drop,alter,reference,trigger,insert,update,delete', 'universe', nil, { if_not_exists = true })

    local model_space = box.schema.space.create(
        os.getenv('CACHE_DATABASE_MODEL_SPACE'),
        { if_not_exists = true }
    )

    model_space:format({
        { name = 'key',   type = 'string' },
        { name = 'value', type = '*' },
    })

    model_space:create_index('primary',
        { type = 'hash', parts = {1, 'string'}, if_not_exists = true }
    )

    local weight_space = box.schema.space.create(
        os.getenv('CACHE_DATABASE_WEIGHT_SPACE'),
        { if_not_exists = true }
    )

    weight_space:format({
        { name = 'key',   type = 'string' },
        { name = 'value', type = '*' },
    })

    weight_space:create_index('primary',
        { type = 'hash', parts = {1, 'string'}, if_not_exists = true }
    )
end)
