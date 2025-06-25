{ pkgs ? import <nixpkgs> { } }:

pkgs.mkShell {
  packages = with pkgs; [
    # Python with exact packages
    (python313.withPackages (ps:
      with ps; [
        brotli
        certifi
        charset-normalizer
        mutagen
        pycryptodomex
        pygame
        requests
        termcolor
        urllib3
        websockets
        yt-dlp
        ytmusicapi
      ]))

    # System dependencies
    ffmpeg
    glib
    zsh
  ];
}
