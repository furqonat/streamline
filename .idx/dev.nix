{ pkgs }: {
  channel = "stable-23.11";
  packages = [
    pkgs.nodejs_20
    pkgs.openjdk17-bootstrap
    pkgs.openssl
    pkgs.docker
    pkgs.go
  ];
  idx.extensions = [

  ];
  idx.previews = {
    previews = {
      web = {
        command = [
          "npx"
          "nx"
          "run"
          "www:dev"
          "--"
          "--port"
          "$PORT"
          "--hostname"
          "0.0.0.0"
        ];
        manager = "web";
      };
      android = {
        command = [ "flutter" "run" "--machine" "-d" "android" "-d" "localhost:5555" ];
        manager = "flutter";
      };
    };
  };
}
