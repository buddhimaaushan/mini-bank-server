CREATE TYPE
    rank AS ENUM (
        'bronze',
        'silver',
        'gold',
        'platinum',
        'garnet',
        'ruby',
        'diamond'
    );

CREATE TYPE
    status AS ENUM (
        'deleted',
        'inactive',
        'holded',
        'active'
    );

CREATE TYPE
    role AS ENUM (
        'admin',
        'manager',
        'bankTeller',
        'customer'
    );