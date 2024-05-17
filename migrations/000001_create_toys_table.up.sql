CREATE TABLE IF NOT EXISTS toys (
    id bigserial PRIMARY KEY,  
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    description TEXT NOT NULL,
    category text NOT NULL,
    rented boolean DEFAULT FALSE,
    price DECIMAL(10, 2) NOT NULL,
    version integer NOT NULL DEFAULT 1
);
ALTER TABLE toys ADD CONSTRAINT toys_category_check CHECK (category IN ('action figures', 'dolls', 'educational', 'games', 'puzzles', 'vehicles'));

ALTER TABLE toys ADD CONSTRAINT toys_price_check CHECK (price >= 0);
