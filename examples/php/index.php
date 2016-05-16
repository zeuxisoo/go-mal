<?php
date_default_timezone_set("Asia/Hong_Kong");

require_once __DIR__.'/vendor/autoload.php';

$transport = Swift_SmtpTransport::newInstance('127.0.0.1', 8053);

$message = Swift_Message::newInstance('Wonderful Subject')
  ->setFrom(['sender@domain.com' => 'Sender Test'])
  ->setTo(['receiver@domain.org' => 'Receiver Other'])
  ->setBody('Here is the message itself');

$mailer = Swift_Mailer::newInstance($transport);
$mailer->send($message);
