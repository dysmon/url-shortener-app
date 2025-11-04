    import React, { useState } from 'react';
    import { shortenUrl } from '../services/api';

    function UrlShortenerForm({ onSuccess, onError }) {
    const [originalUrl, setOriginalUrl] = useState('');
    const [isLoading, setIsLoading] = useState(false);

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (!originalUrl) return;

        setIsLoading(true);
        try {
        const shortCode = await shortenUrl(originalUrl);
        onSuccess(shortCode);
        } catch (err) {
        onError(err.message || 'Failed to shorten URL');
        } finally {
        setIsLoading(false);
        }
    };

    return (
        <form onSubmit={handleSubmit} className="url-form">
        <input
            type="url"
            value={originalUrl}
            onChange={(e) => setOriginalUrl(e.target.value)}
            placeholder="Enter your long URL"
            required
        />
        <button type="submit" disabled={isLoading}>
            {isLoading ? 'Shortening...' : 'Shorten URL'}
        </button>
        </form>
    );
    }

    export default UrlShortenerForm;