export interface ConvertedResult {
    Success?: boolean;
    Error?: null;
    TransformedResults?: { [key: string]: TransformedResult };
}

export interface TransformedResult {
    Format?: string;
    VideoCodec?: string;
    AudioCodec?: string;
    Resolution?: Resolution;
    Quality?: number;
    Data?: string;
    FileExtension?: string
}

export interface Resolution {
    Width?: number;
    Height?: number;
}
