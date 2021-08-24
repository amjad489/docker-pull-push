# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class DockerPullPush < Formula
  desc "Migrate docker images from docker hub to AWS ECR."
  homepage "https://github.com/amjad489/docker-pull-push"
  version "0.3.0"
  license "Apache License 2.0"
  bottle :unneeded

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/amjad489/docker-pull-push/releases/download/v0.3.0/docker-pull-push_0.3.0_Darwin_x86_64.tar.gz"
      sha256 "bbe4ed521d3fc6c42e06cc1e594558c8817f76396e3b4b5cc284a9d0d825331f"
    end
    if Hardware::CPU.arm?
      url "https://github.com/amjad489/docker-pull-push/releases/download/v0.3.0/docker-pull-push_0.3.0_Darwin_arm64.tar.gz"
      sha256 "6b95b97c24be4d9f8170a0b129161a2a41cc226771ba2f0f69fb85f6e0d94f15"
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/amjad489/docker-pull-push/releases/download/v0.3.0/docker-pull-push_0.3.0_Linux_x86_64.tar.gz"
      sha256 "c2cdab3d8275280c6cfd7847d0e800a99157656b92d85bd7501f8c3d508aee9d"
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/amjad489/docker-pull-push/releases/download/v0.3.0/docker-pull-push_0.3.0_Linux_arm64.tar.gz"
      sha256 "53d8aac26e1bd861718bc293ef1c445a2ea5a7a17bc1b0564a33c68cdf8d963b"
    end
  end

  def install
    bin.install "docker-pull-push"
  end
end
